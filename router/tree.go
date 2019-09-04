// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE

package router

import (
	"github.com/hillguo/sanhttp/ctx"
	"net/url"
	"strings"
	"unicode"
)


type MethodTree struct {
	Method string
	ROOT   *Node
}

type MethodTrees []MethodTree

func (trees MethodTrees) Get(method string) *Node {
	for _, tree := range trees {
		if tree.Method == method {
			return tree.ROOT
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

type NodeType uint8

const (
	static NodeType = iota // default
	root
	param
	catchAll
)

type Node struct {
	path      string
	indices   string
	children  []*Node
	Handlers  ctx.HandlersChain
	priority  uint32
	nType     NodeType
	maxParams uint8
	wildChild bool
	FullPath  string
}

// increments priority of the given child and reorders if necessary.
func (n *Node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap Node positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

// addRoute adds a Node with the given handle to the path.
// Not concurrency-safe!
func (n *Node) AddRoute(path string, handlers ctx.HandlersChain) {
	FullPath := path
	n.priority++
	numParams := countParams(path)

	parentFullPathIndex := 0

	// non-empty tree
	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// Update maxParams of the current Node
			if numParams > n.maxParams {
				n.maxParams = numParams
			}

			// Find the longest common prefix.
			// This also implies that the common prefix contains no ':' or '*'
			// since the existing key can't contain those chars.
			i := 0
			max := min(len(path), len(n.path))
			for i < max && path[i] == n.path[i] {
				i++
			}

			// Split edge
			if i < len(n.path) {
				child := Node{
					path:      n.path[i:],
					wildChild: n.wildChild,
					indices:   n.indices,
					children:  n.children,
					Handlers:  n.Handlers,
					priority:  n.priority - 1,
					FullPath:  n.FullPath,
				}

				// Update maxParams (max of all children)
				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}

				n.children = []*Node{&child}
				// []byte for proper unicode char conversion, see #65
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.Handlers = nil
				n.wildChild = false
				n.FullPath = FullPath[:parentFullPathIndex+i]
			}

			// Make new Node a child of this Node
			if i < len(path) {
				path = path[i:]

				if n.wildChild {
					parentFullPathIndex += len(n.path)
					n = n.children[0]
					n.priority++

					// Update maxParams of the child Node
					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--

					// Check if the wildcard matches
					if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
						// check for longer wildcard, e.g. :name and :names
						if len(n.path) >= len(path) || path[len(n.path)] == '/' {
							continue walk
						}
					}

					pathSeg := path
					if n.nType != catchAll {
						pathSeg = strings.SplitN(path, "/", 2)[0]
					}
					prefix := FullPath[:strings.Index(FullPath, pathSeg)] + n.path
					panic("'" + pathSeg +
						"' in new path '" + FullPath +
						"' conflicts with existing wildcard '" + n.path +
						"' in existing prefix '" + prefix +
						"'")
				}

				c := path[0]

				// slash after param
				if n.nType == param && c == '/' && len(n.children) == 1 {
					parentFullPathIndex += len(n.path)
					n = n.children[0]
					n.priority++
					continue walk
				}

				// Check if a child with the next path byte exists
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						parentFullPathIndex += len(n.path)
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}

				// Otherwise insert it
				if c != ':' && c != '*' {
					// []byte for proper unicode char conversion, see #65
					n.indices += string([]byte{c})
					child := &Node{
						maxParams: numParams,
						FullPath:  FullPath,
					}
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				n.insertChild(numParams, path, FullPath, handlers)
				return

			} else if i == len(path) { // Make Node a (in-path) leaf
				if n.Handlers != nil {
					panic("handlers are already registered for path '" + FullPath + "'")
				}
				n.Handlers = handlers
			}
			return
		}
	} else { // Empty tree
		n.insertChild(numParams, path, FullPath, handlers)
		n.nType = root
	}
}

func (n *Node) insertChild(numParams uint8, path string, FullPath string, handlers ctx.HandlersChain) {
	var offset int // already handled bytes of the path

	// find prefix until first wildcard (beginning with ':' or '*')
	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		if c != ':' && c != '*' {
			continue
		}

		// find wildcard end (either '/' or path end)
		end := i + 1
		for end < max && path[end] != '/' {
			switch path[end] {
			// the wildcard name must not contain ':' and '*'
			case ':', '*':
				panic("only one wildcard per path segment is allowed, has: '" +
					path[i:] + "' in path '" + FullPath + "'")
			default:
				end++
			}
		}

		// check if this Node existing children which would be
		// unreachable if we insert the wildcard here
		if len(n.children) > 0 {
			panic("wildcard route '" + path[i:end] +
				"' conflicts with existing children in path '" + FullPath + "'")
		}

		// check if the wildcard has a name
		if end-i < 2 {
			panic("wildcards must be named with a non-empty name in path '" + FullPath + "'")
		}

		if c == ':' { // param
			// split path at the beginning of the wildcard
			if i > 0 {
				n.path = path[offset:i]
				offset = i
			}

			child := &Node{
				nType:     param,
				maxParams: numParams,
				FullPath:  FullPath,
			}
			n.children = []*Node{child}
			n.wildChild = true
			n = child
			n.priority++
			numParams--

			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if end < max {
				n.path = path[offset:end]
				offset = end

				child := &Node{
					maxParams: numParams,
					priority:  1,
					FullPath:  FullPath,
				}
				n.children = []*Node{child}
				n = child
			}

		} else { // catchAll
			if end != max || numParams > 1 {
				panic("catch-all routes are only allowed at the end of the path in path '" + FullPath + "'")
			}

			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				panic("catch-all conflicts with existing handle for the path segment root in path '" + FullPath + "'")
			}

			// currently fixed width 1 for '/'
			i--
			if path[i] != '/' {
				panic("no / before catch-all in path '" + FullPath + "'")
			}

			n.path = path[offset:i]

			// first Node: catchAll Node with empty path
			child := &Node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
				FullPath:  FullPath,
			}
			n.children = []*Node{child}
			n.indices = string(path[i])
			n = child
			n.priority++

			// second Node: Node holding the variable
			child = &Node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				Handlers:  handlers,
				priority:  1,
				FullPath:  FullPath,
			}
			n.children = []*Node{child}

			return
		}
	}

	// insert remaining path part and handle to the leaf
	n.path = path[offset:]
	n.Handlers = handlers
	n.FullPath = FullPath
}

// NodeValue holds return values of (*Node).getValue method
type NodeValue struct {
	Handlers ctx.HandlersChain
	Params   ctx.Params
	tsr      bool
	FullPath string
}

// getValue returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
func (n *Node) GetValue(path string, po ctx.Params, unescape bool) (value NodeValue) {
	value.Params = po
walk: // Outer loop for walking the tree
	for {
		if len(path) > len(n.path) {
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				// If this Node does not have a wildcard (param or catchAll)
				// child,  we can just look up the next child Node and continue
				// to walk down the tree
				if !n.wildChild {
					c := path[0]
					for i := 0; i < len(n.indices); i++ {
						if c == n.indices[i] {
							n = n.children[i]
							continue walk
						}
					}

					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					value.tsr = path == "/" && n.Handlers != nil
					return
				}

				// handle wildcard child
				n = n.children[0]
				switch n.nType {
				case param:
					// find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// save param value
					if cap(value.Params) < int(n.maxParams) {
						value.Params = make(ctx.Params, 0, n.maxParams)
					}
					i := len(value.Params)
					value.Params = value.Params[:i+1] // expand slice within preallocated capacity
					value.Params[i].Key = n.path[1:]
					val := path[:end]
					if unescape {
						var err error
						if value.Params[i].Value, err = url.QueryUnescape(val); err != nil {
							value.Params[i].Value = val // fallback, in case of error
						}
					} else {
						value.Params[i].Value = val
					}

					// we need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						value.tsr = len(path) == end+1
						return
					}

					if value.Handlers = n.Handlers; value.Handlers != nil {
						value.FullPath = n.FullPath
						return
					}
					if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						value.tsr = n.path == "/" && n.Handlers != nil
					}

					return

				case catchAll:
					// save param value
					if cap(value.Params) < int(n.maxParams) {
						value.Params = make(ctx.Params, 0, n.maxParams)
					}
					i := len(value.Params)
					value.Params = value.Params[:i+1] // expand slice within preallocated capacity
					value.Params[i].Key = n.path[2:]
					if unescape {
						var err error
						if value.Params[i].Value, err = url.QueryUnescape(path); err != nil {
							value.Params[i].Value = path // fallback, in case of error
						}
					} else {
						value.Params[i].Value = path
					}

					value.Handlers = n.Handlers
					value.FullPath = n.FullPath
					return

				default:
					panic("invalid Node type")
				}
			}
		} else if path == n.path {
			// We should have reached the Node containing the handle.
			// Check if this Node has a handle registered.
			if value.Handlers = n.Handlers; value.Handlers != nil {
				value.FullPath = n.FullPath
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			for i := 0; i < len(n.indices); i++ {
				if n.indices[i] == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.Handlers != nil) ||
						(n.nType == catchAll && n.children[0].Handlers != nil)
					return
				}
			}

			return
		}

		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		value.tsr = (path == "/") ||
			(len(n.path) == len(path)+1 && n.path[len(path)] == '/' &&
				path == n.path[:len(n.path)-1] && n.Handlers != nil)
		return
	}
}

// findCaseInsensitivePath makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
func (n *Node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (ciPath []byte, found bool) {
	ciPath = make([]byte, 0, len(path)+1) // preallocate enough memory

	// Outer loop for walking the tree
	for len(path) >= len(n.path) && strings.EqualFold(path[:len(n.path)], n.path) {
		path = path[len(n.path):]
		ciPath = append(ciPath, n.path...)

		if len(path) > 0 {
			// If this Node does not have a wildcard (param or catchAll) child,
			// we can just look up the next child Node and continue to walk down
			// the tree
			if !n.wildChild {
				r := unicode.ToLower(rune(path[0]))
				for i, index := range n.indices {
					// must use recursive approach since both index and
					// ToLower(index) could exist. We must check both.
					if r == unicode.ToLower(index) {
						out, found := n.children[i].findCaseInsensitivePath(path, fixTrailingSlash)
						if found {
							return append(ciPath, out...), true
						}
					}
				}

				// Nothing found. We can recommend to redirect to the same URL
				// without a trailing slash if a leaf exists for that path
				found = fixTrailingSlash && path == "/" && n.Handlers != nil
				return
			}

			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				k := 0
				for k < len(path) && path[k] != '/' {
					k++
				}

				// add param value to case insensitive path
				ciPath = append(ciPath, path[:k]...)

				// we need to go deeper!
				if k < len(path) {
					if len(n.children) > 0 {
						path = path[k:]
						n = n.children[0]
						continue
					}

					// ... but we can't
					if fixTrailingSlash && len(path) == k+1 {
						return ciPath, true
					}
					return
				}

				if n.Handlers != nil {
					return ciPath, true
				} else if fixTrailingSlash && len(n.children) == 1 {
					// No handle found. Check if a handle for this path + a
					// trailing slash exists
					n = n.children[0]
					if n.path == "/" && n.Handlers != nil {
						return append(ciPath, '/'), true
					}
				}
				return

			case catchAll:
				return append(ciPath, path...), true

			default:
				panic("invalid Node type")
			}
		} else {
			// We should have reached the Node containing the handle.
			// Check if this Node has a handle registered.
			if n.Handlers != nil {
				return ciPath, true
			}

			// No handle found.
			// Try to fix the path by adding a trailing slash
			if fixTrailingSlash {
				for i := 0; i < len(n.indices); i++ {
					if n.indices[i] == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.Handlers != nil) ||
							(n.nType == catchAll && n.children[0].Handlers != nil) {
							return append(ciPath, '/'), true
						}
						return
					}
				}
			}
			return
		}
	}

	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
	if fixTrailingSlash {
		if path == "/" {
			return ciPath, true
		}
		if len(path)+1 == len(n.path) && n.path[len(path)] == '/' &&
			strings.EqualFold(path, n.path[:len(path)]) &&
			n.Handlers != nil {
			return append(ciPath, n.path...), true
		}
	}
	return
}

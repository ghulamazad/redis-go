package resp

import (
	"sync"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

var Handlers = map[string]func([]Value) Value{
	"PING":    Ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}

func Ping(args []Value) Value {
	if len(args) == 0 {
		return Value{Type: "string", Str: "PONG"}
	}
	return Value{Type: "string", Str: args[0].Bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{Type: "string", Str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk
	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{Type: "null"}
	}

	return Value{Type: "bulk", Bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{Type: "string", Str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{Type: "null"}
	}

	return Value{Type: "bulk", Bulk: value}
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'hsetall' command"}
	}

	hash := args[0].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return Value{Type: "null"}
	}

	var array []Value
	for _, v := range value {
		array = append(array, Value{Type: "bulk", Bulk: v})
	}
	return Value{Type: "array", Array: array}
}

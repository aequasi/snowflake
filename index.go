package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net/http"
	"strconv"
)

var node *snowflake.Node

func H(w http.ResponseWriter, r *http.Request) {
	if node == nil {
		snowflake.Epoch = 1545264000000
		snowflakeNode, err := snowflake.NewNode(1)

		if err != nil {
			panic(err)
		}

		node = snowflakeNode
	}


	count, ok := r.URL.Query()["count"]
	if !ok {
		fmt.Fprintf(w, node.Generate().String())

		return
	}
	parsedCount, err := strconv.ParseInt(count[0], 10, 32)
	if err != nil || parsedCount > 10000 {
		w.WriteHeader(500)
		fmt.Fprintf(w, "`count` must be a number below 10000")

		return
	}

	ids := make([]snowflake.ID, 0)
	for i := 0; i < int(parsedCount); i++ {
		ids = append(ids, node.Generate())
	}

	jsonResponse, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", jsonResponse)
}

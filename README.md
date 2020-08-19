# xd

A stupid `map[string]interface{}` / `[]interface{}` _digger_. Kinda like Ruby's [`Hash#dig`](https://ruby-doc.org/core-2.3.0/Hash.html#method-i-dig).

## Why

Mostly because you'd always need to do some shenanigans with objects that you cannot (or don't want to) unmarshal to typed structures.

And then duck-typing and bounds checking and everything else could give you carpal tunnel syndrome.

Here's just a simple example:

``` go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tzvetkoff-go/xd"
)

func main() {
	jsonString := `{
		"data": [
			{
				"type": "Article",
				"id": 1,
				"attributes": {
					"title": "Who said JSON parsing should be painful?"
				}
			}
		]
	}`

	m := map[string]interface{}{}
	json.Unmarshal([]byte(jsonString), &m)

	//
	// Because this:
	//

	if mData, ok := m["data"]; ok {
		if mData, ok := mData.([]interface{}); ok {
			if len(mData) > 0 {
				mData0 := mData[0]
				if mData0, ok := mData0.(map[string]interface{}); ok {
					if mData0Attributes, ok := mData0["attributes"]; ok {
						if mData0Attributes, ok := mData0Attributes.(map[string]interface{}); ok {
							if mData0AttributesTitle, ok := mData0Attributes["title"]; ok {
								fmt.Println(mData0AttributesTitle)
							}
						}
					}
				}
			}
		}
	}

	//
	// Could've been this:
	//

	val := xd.Dig(m, "data[0].attributes.title")
	fmt.Println(val)
}
```

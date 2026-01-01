package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type SampleUser struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" jsonschema:"title=the name,description=The name of a friend,example=joe,example=lucy,default=alex"`
	Friends     []int          `json:"friends,omitempty" jsonschema_description:"The list of IDs, omitted when empty"`
	Tags        map[string]any `json:"tags,omitempty" jsonschema_extras:"a=b,foo=bar,foo=bar1"`
	BirthDate   time.Time      `json:"birth_date,omitempty" jsonschema:"oneof_required=date"`
	YearOfBirth string         `json:"year_of_birth,omitempty" jsonschema:"oneof_required=year"`
	Metadata    any            `json:"metadata,omitempty" jsonschema:"oneof_type=string;array"`
	FavColor    string         `json:"fav_color,omitempty" jsonschema:"enum=red,enum=green,enum=blue"`
}

func main() {
	//获取当前路径
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f := filepath.Join(wd, "base/jsonschema/test_user.json")

	expectedJSON, err := os.ReadFile(f)
	fmt.Println(string(expectedJSON))
}

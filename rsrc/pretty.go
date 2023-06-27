package rsrc

import (
	"fmt"
	"sort"
)

type idFile struct {
	id   uint16
	path string
}

type byId []idFile

func (id byId) Len() int {
	return len(id)
}

func (id byId) Less(i, j int) bool {
	return id[i].id < id[j].id
}

func (id byId) Swap(i, j int) {
	id[i], id[j] = id[j], id[i]
}

func PrintIds(pack string, ids map[string]uint16) {
	var idList byId
	for path, id := range ids {
		idList = append(idList, idFile{id, path})
	}
	sort.Sort(idList)
	fmt.Printf(`// generated by github.com/dentalwings/rsrc

package %s

// exeIDs maps the original file names to their resource IDs when embedded in
// the executable.
var exeIDs = map[string]uint16{
`, pack)

	for _, id := range idList {
		fmt.Printf("\n\t\"%s\": %v,", id.path, id.id)
	}

	if len(ids) > 0 {
		fmt.Println()
	}

	fmt.Println("}")
}

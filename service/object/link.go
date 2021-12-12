package object

import (
	"sort"

	"github.com/crackeer/goweb/model"
)

// GetAllLinkList
//  @return []map
func GetAllLinkList() []map[string]interface{} {
	objects, _ := model.GetAll(model.TypeLink)

	mapData := map[string][]map[string]interface{}{}
	for _, tmp := range objects {
		if l, ok := mapData[tmp.Tag]; ok {
			l = append(l, tmp.ToMap())
			mapData[tmp.Tag] = l
			continue
		}

		mapData[tmp.Tag] = []map[string]interface{}{tmp.ToMap()}
	}

	retData := []map[string]interface{}{}
	for tag, l := range mapData {
		retData = append(retData, map[string]interface{}{
			"tag":  tag,
			"list": l,
		})
	}

	sort.Slice(retData, func(i, j int) bool {
		tag1, _ := retData[i]["tag"].(string)
		tag2, _ := retData[j]["tag"].(string)
		return tag1 > tag2
	})
	return retData
}

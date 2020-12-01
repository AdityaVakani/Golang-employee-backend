package models

type MenuItem struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Link      string     `json:"link"`
	Component string     `json:"component"`
	Variant   string     `json:"variant"`
	Icon      string     `json:"icon"`
	ParentID  int        `json:"parentId"`
	SubMenu   []MenuItem `json:"subMenu"`
}

func MenuHirearchy(allMenu []MenuItem) MenuItem {
	for i := len(allMenu) - 1; i >= 0; i-- {
		bottomPID := allMenu[i].ParentID
		if bottomPID != 0 {
			for j := range allMenu {
				row := &allMenu[j]
				if row.ID == bottomPID {
					row.SubMenu = append(row.SubMenu, allMenu[i])

				}
			}
		}
	}
	return allMenu[0]
}

package models

type MenuItem struct {
	ID        int
	Name      string
	Link      string
	Component string
	Variant   string
	Icon      string
	ParentID  int
	SubMenu   []MenuItem
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

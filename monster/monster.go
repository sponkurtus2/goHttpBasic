package monster

import "strconv"

type Monster struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func monsters() []Monster {
	return []Monster{
		{
			ID:   1,
			Name: "Carlos",
		},
		{
			ID:   2,
			Name: "Fernanda",
		},
		{
			ID:   3,
			Name: "Shock",
		},
	}
}

func loadMonstersById() map[string]Monster {
	monsters := monsters()

	res := make(map[string]Monster, len(monsters))

	for _, x := range monsters {
		res[strconv.Itoa(int(x.ID))] = x
		// res[x.name] = x
	}
	return res
}

/*
func loadMonstersByName() map[string]Monster {
	monsters := monsters()

	res := make(map[string]Monster, len(monsters))

	for _, x := range monsters {
		res[x.Name] = x
	}

	return res
}
*/

package seed

import "gorm.io/gorm"

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				var err error
				for i := 0; i < 10; i++ {
					err = CreateUser()
					if err != nil {
						break
					}
				}
				return err
			},
		},
		// {
		// 	Name: "CreateProduct",
		// 	Run: func(db *gorm.DB) error {
		// 		var err error
		// 		for i := 0; i < 10; i++ {
		// 			err = CreateProduct(db)
		// 			if err != nil {
		// 				break
		// 			}
		// 		}
		// 		return err
		// 	},
		// },
	}
}
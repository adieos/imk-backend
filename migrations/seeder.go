package migrations

import (
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	// WARNING: WILL produce duplicate (bc no id in .json)

	/*
		if err := seeds.ProdCodeSeeder(db); err != nil {
			return err
		}

		if err := seeds.ProdFacultySeeder(db); err != nil {
			return err
		}

		if err := seeds.ProdMajorSeeder(db); err != nil {
			return err
		}

		if err := seeds.ProdProdiSeeder(db); err != nil {
			return err
		}

		if err := seeds.ProdProductSeeder(db); err != nil {
			return err
		}
	*/

	// BELOW NOT FOR PROD

	// if err := seeds.ListFacultySeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListMajorSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListProvinceSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListUserSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListOCSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListProductSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListFordaSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListIljFacultySeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListTransactionSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListIljMajorSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListIljQuestionSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListFordaSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListWSNCodeSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListWSNCodeSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListWSNSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListStaffSeeder(db); err != nil {
	// 	return err
	// }

	// if err := seeds.ListFordaLPSeeder(db); err != nil {
	// 	return err
	// }

	return nil
}

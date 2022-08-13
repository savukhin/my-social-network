package models

func UpdateUser(id int, fields map[string]string) error {
	columns_queue := ""
	for key, value := range fields {
		columns_queue += key + " = " + value
	}
	// sql := fmt.Sprintf(`
	// 	UPDATE
	// 		users
	// 	SET
	// 		column1 = value1,
	// 		column2 = value2
	// 	WHERE
	// 		id = %d;
	// `)
	return nil
}

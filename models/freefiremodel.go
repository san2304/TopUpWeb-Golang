package models

import (
	"topup-game/config"
	"topup-game/entities"
)

func GetDiamondOptions() ([]entities.DiamondOption, error) {
	rows, err := config.DB.Query("SELECT product_id, product_name, value, harga, stock FROM freefire")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []entities.DiamondOption
	for rows.Next() {
		var option entities.DiamondOption
		if err := rows.Scan(&option.ProductID, &option.ProductName, &option.Value, &option.Harga, &option.Stock); err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	return options, nil
}

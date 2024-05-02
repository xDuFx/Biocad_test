package repository

import (
	"context"
	"testB/pkg/models"
)


func (repo *PGRepo) NewData(data []models.DataTcv) error {
	for i := 0; i < len(data); i++ {
		_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO DataFiles(n, mqtt, invid, unit_guid, msg_id, text_f, context, class_f, level_f, area_f, addr, block, type_f, bit_f, invert_bit)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`,
			data[i].N,
			data[i].Mqtt,
			data[i].Invid,
			data[i].Unit_guid,
			data[i].Msg_id,
			data[i].Text,
			data[i].Context,
			data[i].Class,
			data[i].Level,
			data[i].Area,
			data[i].Addr,
			data[i].Block,
			data[i].TypE,
			data[i].Bit,
			data[i].Invert_bit,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *PGRepo) NewFilename(filename string) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO listFiles(filename)
    VALUES ($1);`,
		filename,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PGRepo) GetFilenames() ([]string, error) {
	rows, err := repo.pool.Query(context.Background(), `
	SELECT * FROM listFiles;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(
			&name,
		)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func (repo *PGRepo) Pagin(page, limit int) ([]models.DataTcv, error) {
	page = (page - 1) * limit
	_, err := repo.pool.Exec(context.Background(), `
	UPDATE DataFiles
	SET mqtt = 'empty' , block = 'empty', type_f = 'empty', bit_f = 'empty', invert_bit = 'empty'
	WHERE n  between 1 and 1000; `,
	)
	if err != nil {
		return nil, err
	}
	rows, err := repo.pool.Query(context.Background(), `
	SELECT n, mqtt, invid, unit_guid, msg_id, text_f, context, class_f, level_f, area_f, addr, block, type_f, bit_f, invert_bit 
	FROM DataFiles WHERE n > $2 
	ORDER BY n ASC
	LIMIT $1;`, limit, page,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var datas []models.DataTcv
	for rows.Next() {
		var data models.DataTcv
		err := rows.Scan(
			&data.N,
			&data.Mqtt,
			&data.Invid,
			&data.Unit_guid,
			&data.Msg_id,
			&data.Text,
			&data.Context,
			&data.Class,
			&data.Level,
			&data.Area,
			&data.Addr,
			&data.Block,
			&data.TypE,
			&data.Bit,
			&data.Invert_bit,
		)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (repo *PGRepo) Errors(er string) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO errors(error)
    VALUES ($1);`,
		er,
	)
	if err != nil {
		return err
	}
	return nil
}

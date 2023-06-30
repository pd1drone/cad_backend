package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func DebugAnalogVoltageValue(db sqlx.Ext, voltagedata float64) error {
	_, err := db.Exec(`INSERT INTO debug_glocuse_analog(
		VoltageData,
		Timestamp
	)
	Values(?,?)`,
		voltagedata,
		time.Now().Unix(),
	)
	if err != nil {
		return err
	}

	return nil
}

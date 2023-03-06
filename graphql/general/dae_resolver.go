/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, v2rayA Organization <team@v2raya.org>
 */

package general

import (
	"context"
	"github.com/v2rayA/dae-wing/db"
)

type DaeResolver struct {
	Ctx context.Context
}

func (r *DaeResolver) Running() (bool, error) {
	var m db.System
	q := db.DB(r.Ctx).Model(&db.System{}).FirstOrCreate(&m)
	if q.Error != nil {
		return false, q.Error
	}
	return m.Running, nil
}

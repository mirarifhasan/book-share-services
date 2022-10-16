/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `json:"name"`
}

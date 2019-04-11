/* ######################################################################
# File Name: libs/utils/types.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-03-27 13:32:50
####################################################################### */
package utils

// http
type (
	MultipartField struct {
		IsFile    bool
		Fieldname string
		Value     []byte
		Filename  string
	}
)

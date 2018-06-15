<?php
/**
 * Generate the GroupAlias.go
 */

echo <<<EOF
package router

import (

)

/*
Generated with "Generate_GroupAlias.php"
*/


EOF;

$aliases = ["Get", "Post", "Put", "Delete", "Patch", "CLI", "WebSocket"];

foreach ($aliases as $alias) {
    echo <<<EOF
func (this *Group) {$alias}(uri string, controller interface{}) *Rule {
	return this.On("{$alias}", uri, controller)
}


EOF;
}
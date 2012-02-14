<?php
require_once 'Console/CommandLine.php';

$parser = new Console_CommandLine();
$parser->description = 'Generates Go source from input text file and generates array of data.';
$parser->version = '1.5.0';
$parser->addOption('filename', array(
    'short_name'  => '-f',
    'long_name'   => '--file',
    'description' => 'read from from file or url (defaults to stdin)',
    'action'      => 'StoreString'
));
$parser->addOption('output', array(
    'short_name'  => '-o',
    'long_name'   => '--output',
    'description' => "output to file or url (defaults to stdout)",
    'action'      => 'StoreString'
));
$parser->addOption('package', array(
    'short_name'  => '-p',
    'long_name'   => '--package',
    'description' => 'package name (defaults to main)',
    'action'      => 'StoreString'
));
$parser->addOption('name', array(
    'short_name'  => '-n',
    'long_name'   => '--name',
    'description' => 'variable name (defaults to MyVar)',
    'action'      => 'StoreString'
));
$parser->addOption('type', array(
    'short_name'  => '-t',
    'long_name'   => '--type',
    'description' => 'the output type; map(map[string]bool) or array([]string) (defaults to array)',
    'action'      => 'StoreString'
));
try {
    $result = $parser->parse();
    // do something with the result object
    //print_r($result->options);
    //print_r($result->args);
} catch (Exception $exc) {
    $parser->displayError($exc->getMessage());
}

$input = (is_null($result->options['filename']))?'php://stdin':$result->options['filename'];
$lines = file($input, FILE_IGNORE_NEW_LINES|FILE_SKIP_EMPTY_LINES);

// Escape quote chars
$lines = str_replace('"', '\"', $lines);

$output = (is_null($result->options['output']))?'php://stdout':$result->options['output'];
$out = fopen($output, 'w+');

$packageName = (is_null($result->options['package']))?'main':$result->options['package'];
$arrayName = (is_null($result->options['name']))?'MyArray':$result->options['name'];
$type = (is_null($result->options['type']))?'array':$result->options['type'];

    switch($type) {
        case 'array':
            $head = <<<head
package %s

var %s = []string {
head;
            $foot = <<<foot
}
foot;
            $linePrefix = '"';
            $lineSuffix = "\",\n";
            break;

        case 'map':
            $head = <<<head
package %s

var %s = map[string]bool {
head;
            $foot = <<<foot
}
foot;
            $linePrefix = '"';
            $lineSuffix = "\": true,\n";
            break;

        default:
            echo "Unrecognised type\n";
            exit(1);
    }


function modifyLines($line) {
    global $linePrefix, $lineSuffix;
    if(strlen($line)==0)
        return $line;

    return $linePrefix . $line . $lineSuffix;
}

$lines = array_map("modifyLines", $lines);

fwrite($out, sprintf($head, $packageName, $arrayName));
fwrite($out, implode("", $lines));
fwrite($out, $foot);
fclose($out);



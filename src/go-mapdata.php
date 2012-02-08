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
    'description' => 'array name (defaults to MyArray)',
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

$output = (is_null($result->options['output']))?'php://stdout':$result->options['output'];
$out = fopen($output, 'w+');

$packageName = (is_null($result->options['package']))?'main':$result->options['package'];
$arrayName = (is_null($result->options['name']))?'MyArray':$result->options['name'];

$head = <<<head
package %s

var %s = []string {
head;

$foot = <<<foot
}
foot;

fwrite($out, sprintf($head, $packageName, $arrayName));
fwrite($out, sprintf("\n\"%s\"\n", implode("\",\n\"", $lines)));
fwrite($out, $foot);
fclose($out);



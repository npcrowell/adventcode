OUTFILE=advent.elf
ARGS="$@"
ADDITIONALARGS="< data/d02.txt"
COMPILE="go build -o $OUTFILE ."
EXECUTE="./$OUTFILE $ARGS $ADDITIONALARGS"
TIDY="go mod tidy"

echo "TIDYING: $TIDY"
$TIDY
rc=$?

if [ $rc -ne 0 ]; then
exit $rc
fi
echo "TIDYING SUCCESSFULL"

echo "COMPILING: $COMPILE"
$COMPILE
rc=$?

if [ $rc -ne 0 ]; then
exit $rc
fi
echo "COMPILATION SUCCESSFUL"

echo "EXECUTING: $EXECUTE"
echo "============================="
$EXECUTE

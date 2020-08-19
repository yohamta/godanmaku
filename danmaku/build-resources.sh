INPUT_FILE=
OUTPUT_FILE=
PACKAGE_NAME=
VARIABLE_NAME=

DIRECTORY=resources/images
PACKAGE_NAME=images

for f in `ls $DIRECTORY`; do
  echo "processing $f"
  INPUT_FILE=$DIRECTORY/$f
  OUTPUT_FILE=$DIRECTORY/${f%.*}.go
  VARIABLE_NAME=`echo ${f%.*} | tr '[:lower:]' '[:upper:]'`
  file2byteslice -input $INPUT_FILE -output $OUTPUT_FILE -package $PACKAGE_NAME -var $VARIABLE_NAME
done
#! /bin/sh

export ROOT=$PWD
export SRC_DIR=src/
export TEMPLATES_DIR=templates/
export DALI_SERVER_DIR=docker/dali-server/
export MONGO_DIR=docker/mongodb/

# build dali-server
cp $SRC_DIR/* $DALI_SERVER_DIR$SRC_DIR
cp $TEMPLATES_DIR/* $DALI_SERVER_DIR$TEMPLATES_DIR

cd $DALI_SERVER_DIR
docker build -t dali-server .

cd $ROOT

rm $DALI_SERVER_DIR$SRC_DIR/*.go
rm $DALI_SERVER_DIR$TEMPLATES_DIR/*.jpg
rm $DALI_SERVER_DIR$TEMPLATES_DIR/*.html
rm $DALI_SERVER_DIR$TEMPLATES_DIR/*.css

# build mongo image for dali
cd $MONGO_DIR
docker build -t dali-mongo .



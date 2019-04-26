#!/usr/bin/env bash

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then 
  brew install s3cmd && brew link s3cmd
  brew install gtk+ && brew link gtk+
  brew install gtk+3 && brew link gtk+3
  brew install opencv && brew link opencv
else
  sudo apt-get update
  sudo apt-get install -y - libgmp-dev build-essential cmake git gtk2.0 gtk+3.0 libgtk2.0-dev libgtk-3-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libjasper-dev libdc1394-22-dev python-setuptools
  wget http://netix.dl.sourceforge.net/project/s3tools/s3cmd/2.0.2/s3cmd-2.0.2.tar.gz
  tar xvfz s3cmd-2.0.2.tar.gz
  cd s3cmd-2.0.2
  sudo python setup.py install
  cd ..
fi
#!/usr/bin/env python3

import argparse
import subprocess
import os
import sys
import platform
import shutil

SENDER = 'sender.go'
RECEIVER = 'receiver.go'
BUILD_DIR = 'build'
PROTO_SRC = 'protobuf'
PROTO_BUILD_DIR = PROTO_SRC + '/build'

WINBUILD = platform.system() == "Windows"


def verify_deps():
    def is_exe(exe_name):
        for path in os.environ["PATH"].split(os.pathsep):
            exe_fp = os.path.join(path, exe_name) if not WINBUILD else os.path.join(path, exe_name + '.exe')
            if os.path.isfile(exe_fp) and os.access(exe_fp, os.X_OK):
                return True
        return False

    if not is_exe('go'):
        print("Unable to build. Problem locating or executing the 'go' binary")
        return False
    if not is_exe('protoc'):
        print("Unable to build. Problem locating or executing the 'protoc' binary")
        return False
    if not is_exe('protoc-gen-go'):
        print("Unable to build. Problem locating or executing the 'protoc-gen-go' binary")
        return False
    return True


def build_protobufs():
    if not os.path.isdir(PROTO_BUILD_DIR):
        os.mkdir(PROTO_BUILD_DIR)

    for filename in os.listdir(PROTO_SRC):
        if filename.endswith('.proto'):
            cmd = 'protoc -I={} --go_out={} {}'.format(PROTO_SRC, PROTO_BUILD_DIR, os.path.join(PROTO_SRC, filename))
            print(cmd)
            subprocess.run(cmd.split(), check=True, stdout=sys.stdout, stderr=sys.stderr)


def build_go():
    cmd = "go build -v -o {}/Sender{} {}".format(BUILD_DIR, "" if not WINBUILD else ".exe", SENDER)
    print(cmd)
    subprocess.run(cmd.split(), stdout=sys.stdout, stderr=sys.stderr, check=True)

    cmd = "go build -v -o {}/Receiver{} {}".format(BUILD_DIR, "" if not WINBUILD else ".exe", RECEIVER)
    print(cmd)
    subprocess.run(cmd.split(), stdout=sys.stdout, stderr=sys.stderr, check=True)


def clean():
    print("Cleaning...")
    try:
        shutil.rmtree(BUILD_DIR)
    except FileNotFoundError:
        pass
    try:
        shutil.rmtree(PROTO_BUILD_DIR)
    except FileNotFoundError:
        pass


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='This script builds the protobuf-networking project')
    parser.add_argument('--clean', '-c', action='store_true', help='Perform a clean build', default=False)
    args = parser.parse_args()

    if args.clean:
        clean()

    if not verify_deps():
        exit()

    build_protobufs()
    build_go()

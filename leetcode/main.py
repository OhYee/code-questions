#!/bin/python3
import json
import sys
import os
import subprocess

PWD = os.path.dirname(os.path.abspath(__file__))

def color(num, *args, **kwargs):
    end = kwargs.get("end", "\n")
    if "end" in kwargs:
        del kwargs["end"]
    print("\033[%sm" % (str(num)), end="")
    print(*args, **kwargs, end="")
    print("\033[0m", end=end)


class Context:
    def __init__(self, _id, name, tests=[], language="python", debug=False, verbose=False):
        self.id = _id
        self.language = language
        self.debug = debug
        self.name = name
        self.tests = tests
        self.verbose = verbose

    def __repr__(self):
        return "id={} name={} language={} debug={} verbose={}\n{}".format(
            self.id, self.name, self.language,
            self.debug, self.verbose, self.tests
        )


def call(*args, **kwargs):
    proc = subprocess.Popen(
        " ".join(args),
        shell=True,
        stdout=subprocess.PIPE,
        stdin=subprocess.PIPE,
        bufsize=0
    )
    try:
        outs, errs = proc.communicate(
            timeout=15,
            input=bytes(
                json.dumps(kwargs["input"]),
                encoding="utf-8",
            ) if "input" in kwargs else None,
        )
        outs = str(outs, encoding="utf-8")
    except Exception as e:
        outs = ""
        color(1, e)
        color(1, errs)
    finally:
        proc.kill()
    try:
        res = json.loads(outs)["out"]
    except:
        res = outs
    return  res
from typing import *

{code}

if __name__ == "__main__":
    solution = Solution()
    tests = {tests}
    total = len(tests)
    passed = True
    for idx, t in enumerate(tests):
        res = solution.{name}(t["in"])
        if res == t["out"]:
            pass
        else:
            passed = False
            print("\033[91mError at", idx+1, "test\033[0m")
            print(" ", "\033[1minput\033[0m", t["in"])
            print(" ", "\033[1mgot  \033[0m", res)
            print(" ", "\033[1mwant \033[0m", t["out"])

    if passed:
        print("\033[92mAll %d testcases passed\033[0m" % total)
'''


def runPython3(_id, name, tests):
    folder = "%s/%s" % (PWD, _id)
    source = os.path.join(folder, "%s.py" % _id)
    tmpFolder = os.path.join(folder, "dist")
    tmpFile = os.path.join(tmpFolder, "%s.py" % _id)
    with open(source) as f:
        code = f.read()
    if not os.path.exists(tmpFolder):
        os.mkdir(tmpFolder)
    with open(tmpFile, "w") as f:
        f.write(py3template.format(code=code, name=name, tests=tests))

    res = call("python3", tmpFile)
    print(res)


if __name__ == "__main__":
    _id = sys.argv[1]
    with open("./%s/problem.json" % _id) as f:
        problem = json.loads(f.read())
    name = problem["name"]
    tests = problem["samples"]
    runPython3(_id, name, tests)

#!/usr/bin/env python3.8

from google.protobuf import descriptor

import workshop.v1.annotation_pb2 as annopb
import workshop.v1.workshop_pb2 as pb


def apply(m):
    def hide(fd):
        return fd.GetOptions().Extensions[annopb.my_annotation].hide

    for fd in m.DESCRIPTOR.fields:
        if hide(fd):
            m.ClearField(fd.name)
            continue
        if fd.type == descriptor.FieldDescriptor.TYPE_MESSAGE:
            v = getattr(m, fd.name)
            apply(v)


def run():
    apple = pb.Apple()
    apple.brand = "granny-smith"
    apple.age = 42
    apple.skin.colour = "green"
    apple.skin.blemishes = 3

    print("==== before ====")
    print(apple)
    print("==== apply annotations ====")
    apply(apple)
    print(apple)


if __name__ == "__main__":
    run()

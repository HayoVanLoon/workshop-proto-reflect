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
            # go deeper
            v = getattr(m, fd.name)
            apply(v)


def run():
    apple = pb.Apple(
        brand="granny-smith",
        age=42,
        skin=pb.Apple.Skin(
            colour="green",
            blemishes=3
        ),
    )

    print("==== before ====")
    print(apple)
    print("==== after apply ====")
    apply(apple)
    print(apple)


if __name__ == "__main__":
    run()

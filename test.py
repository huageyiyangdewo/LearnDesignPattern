import threading


class SingletonType(type):
    _instance_lock = threading.Lock()
    def __call__(cls, *args, **kwargs):
        # if not hasattr(cls, "_instance"):
        with SingletonType._instance_lock:
            if not hasattr(cls, "_instance"):
                cls._instance = super(SingletonType,cls).__call__(*args, **kwargs)
        return cls._instance


class Foo(metaclass=SingletonType):
    def __init__(self,name):
        import time
        time.sleep(1)
        self.name = name





if __name__ == "__main__":

    obj1 = Foo('name')
    obj2 = Foo('name')
    print(obj1, obj2)
    print(id(obj1))

    import threading
    import time

    def task(arg):
        obj = Foo(arg)
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()

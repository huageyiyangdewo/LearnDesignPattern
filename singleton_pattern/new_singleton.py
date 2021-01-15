import threading


# 方式四：基于 __new__, 线程不安全
class Singleton(object):

    def __init__(self):
        pass

    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            cls._instance = object.__new__(cls)
        return cls._instance


# 基于 __new__, 加锁安全
class SingletonLock(object):

    _instance_lock = threading.Lock()

    def __init__(self):
        pass

    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with cls._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = object.__new__(cls)
        return cls._instance


if __name__ == "__main__":

    def task(arg):
        obj = SingletonLock()
        print(id(obj))


    for i in range(10):
        t = threading.Thread(target=task, args=[i, ])
        t.start()

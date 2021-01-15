class ModelSingleton:
    # 使用模块实现单例

    db_port = 3306

    # 以下是多线程中模拟耗时操作
    # import time
    # time.sleep(1)


ServerConfigSingleton = ModelSingleton()

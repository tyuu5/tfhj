def open_file_A():
    OP_A = open('./config/A.txt', 'r')
    HQ_A = OP_A.read()
    return HQ_A.split()
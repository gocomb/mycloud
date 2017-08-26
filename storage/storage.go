package storage

/*
存储相关操作

初步是想通过ceph的rbd来作为mycloud的后端存储，api-server接收到存储创建的请求，调用这里的方法，利用ceph-common或者calamari的接口创建
相应的rbd块存储，供直接挂载、pv等的使用

*/
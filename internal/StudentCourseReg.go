package internal

/*
student course registration system:  create a course storage with add_course functions etc, no duplicates....(course id, name, credit)
create a student course reg storage system, student can only register courses in the course storage, can't register more than 24 credits (student id, reg course id, credit)
inputs and outputs are all strings, but credit is integer, so need conversion all the time.
需要注意的是course有专门的add_course来添加, 但student没有. 所以在register的时候需要check给定student是否存在, 没存在创建一个, 而course不存在的话需要判定注册失败.
求加米加米

*/

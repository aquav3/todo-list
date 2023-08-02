package tasks

// Gets a pcg_hash
func pcg_hash(input *uint64){
     var state uint64
     var word uint64
     state = *input * 747796405 + 2891336453
     word = ((state >> ((state >> 28) + 4)) ^ state) * 277803737
     *input = (word >> 22) ^ word
}

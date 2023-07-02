# AES Merkle Tree Hash

I was nerd sniped by someone about making a cryptographically safe AES based hash function.

To be clear, DO NOT USE THIS FOR ANYTHING IMPORTANT, it has absolutely no cryptanalysis backing it up.

You should be probably not use it for anything at all either given it's not particularly efficient as it only half the size of the working state on each merkle tree depth.

The only design goal is that I'm not allowed to do any operations on the hash state that is not AES.

This implementation is not efficient by any means either.

The only good points I am aware of are:
- can be incrementally verified
- supports incremental random order verification (AKA provable `.ReadAt`)
- the size is provable (note you can forge a hash that claim any size, however for a properly constructed hash, assuming this is not broken, you can't forge a verification input that claims an other size, similar to blake3)
- Supports parallel hashing (but given it's not optimized for speed to begin with who really cares ?)

## Entry point

It works by first taking the input buffer and zero padding it to a multiple of 16 bytes, we then append at the end a uint128 little endian encoded number which is the length of the input buffer.

This is hashed using the greedy combine strategy to get a 16bytes state which is the final value.

## Greedy combine

If you have 1 single block, encrypt 16 zeros using the block as the key and return that.

Find the biggest power of two blocks in passed in, if this is equal to the total (this happens if you are passed a power of amount two blocks) then take half that.

Split the input buffer in an lhs that is the number computed right above, and rhs the other one.

Run greedy hash on both lhs and rhs, then encrypt the hash of rhs with the key of lhs return that value.
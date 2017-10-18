#ifndef ___VECTOR_H___
#define ___VECTOR_H___

#include<inttypes.h>
#include "util.h"

struct Vector {
	uint32_t cap;
	uint32_t count;
	char data[0];
};

struct Vector* VectorCreate(COMPARE cmp, DESTROY dest, uint32_t cap);
void VectorDestroy(struct Vector* vec);
int VectorInsert(struct Vector* vec, void* data);
int VectorInsertSorted(struct Vector* vec, void* data);
int VectorInsertFront(struct Vector* vec, void* data);
int VectorInsertLast(struct Vector* vec, void* data);
void* VectorDeleteFirst(struct Vector* vec);
void* VectorDeleteLast(struct Vector* vec);
void* VectorDelete(struct Vector* vec, void* key);
void* VectorGetFirst(struct Vector* vec);
void* VectorGetLast(struct Vector* vec);
void* VectorGet(struct Vector* vec, void* key);
void* VectorForEach(struct Vector* vec, OPERATE op);
void* VectorSort(struct Vector* vec);
#endif /* ___VECTOR_H___ */

#include "pppp.h"
#define DEBUG
#include "fault.h"
#include "list.h"
#include "leak_detector.h"
#include "util.h"
#include <stdio.h>
#include "dlist.h"

int compare(void* a, void* b) {
	size_t c = (size_t)a;
	size_t d = (size_t)b;

	return c - d;
}

int destroy(void* a) {
	return 0;
}

int print(void* a) {
	size_t c = (size_t)a;

	printf("(%lu)\n", c);

	return 0;
}

void ListFunctionTest() {
	struct List *list = NULL;
	list = ListCreate(compare, destroy);
	if (list == NULL) {
		printf("Cannot create new list\n");
		return;
	}

	ListInsertFront(list, (void*)100);
	ListInsertFront(list, (void*)10);
	ListInsertFront(list, (void*)101);
	ListInsertFront(list, (void*)104);
	ListInsertFront(list, (void*)100);
	ListInsertFront(list, (void*)100);
	ListInsertFront(list, (void*)109);
	ListInsertFront(list, (void*)10);
	ListInsertFront(list, (void*)100);
	ListInsertFront(list, (void*)100);

	ListForEach(list, print);
	ListDestroy(list);

	list = ListCreate(compare, destroy);
	if (list == NULL) {
		printf("Cannot create new list\n");
		return;
	}

	ListInsertEnd(list, (void*)100);
	ListInsertEnd(list, (void*)10);
	ListInsertEnd(list, (void*)101);
	ListInsertEnd(list, (void*)104);
	ListInsertEnd(list, (void*)100);
	ListInsertEnd(list, (void*)100);
	ListInsertEnd(list, (void*)109);
	ListInsertEnd(list, (void*)10);
	ListInsertEnd(list, (void*)100);
	ListInsertEnd(list, (void*)100);

	ListForEach(list, print);
	ListDeleteFirst(list);
	ListForEach(list, print);
	ListDeleteLast(list);
	ListForEach(list, print);

	ListDeleteAll(list, (void*)100);
	ListForEach(list, print);
	ListReverse(list);
	ListForEach(list, print);
	ListFlush(list);
	ListForEach(list, print);

	ListInsertSorted(list, (void*)1);
	ListInsertSorted(list, (void*)2);
	ListInsertSorted(list, (void*)3);
	ListInsertSorted(list, (void*)10);
	ListInsertSorted(list, (void*)101);
	ListInsertSorted(list, (void*)104);
	ListInsertSorted(list, (void*)109);
	ListInsertSorted(list, (void*)10);
	ListInsertSorted(list, (void*)11);
	ListInsertSorted(list, (void*)12);
	ListInsertSorted(list, (void*)13);
	ListInsertSorted(list, (void*)14);
	ListInsertSorted(list, (void*)15);
	ListInsertSorted(list, (void*)16);
	ListInsertSorted(list, (void*)17);
	ListInsertSorted(list, (void*)100);
	ListInsertSorted(list, (void*)1000);
	ListInsertSorted(list, (void*)1001);
	ListInsertSorted(list, (void*)1002);
	ListInsertSorted(list, (void*)1003);
	ListInsertSorted(list, (void*)1004);
	ListInsertSorted(list, (void*)20);
	ListForEach(list, print);
}

void DListFunctionTest() {
	struct DList *list = NULL;
	list = DListCreate(compare, destroy);
	ASSERT(list != NULL, , "%s", "Cannot create new list for test");

	DListInsertFront(list, (void*)1000);
	DListInsertFront(list, (void*)10);
	DListInsertFront(list, (void*)8000);
	DListInsertFront(list, (void*)2000);
	DListInsertFront(list, (void*)1);
	DListInsertFront(list, (void*)13);
	DListInsertFront(list, (void*)13);
	DListInsertFront(list, (void*)53);
	DListInsertFront(list, (void*)33);
	DListInsertFront(list, (void*)23);
	DListForEach(list, print);
	DListReverse(list);
	DListForEach(list, print);
	DListReverse(list);
	DListForEach(list, print);
	DListDeleteFirst(list);
	DListForEach(list, print);
	DListDeleteFirst(list);
	DListForEach(list, print);

	DListDeleteLast(list);
	DListForEach(list, print);
	DListDelete(list, (void*)13);
	DListForEach(list, print);
	DListDestroy(list);

	list = DListCreate(compare, destroy);
	ASSERT(list != NULL, , "%s", "Cannot create new list for test");

	DListInsertSorted(list, (void*)1000);
	DListInsertSorted(list, (void*)10);
	DListInsertSorted(list, (void*)8000);
	DListInsertSorted(list, (void*)2000);
	DListInsertSorted(list, (void*)1);
	DListInsertSorted(list, (void*)13);
	DListInsertSorted(list, (void*)13);
	DListInsertSorted(list, (void*)53);
	DListInsertSorted(list, (void*)33);
	DListInsertSorted(list, (void*)23);
	DListForEach(list, print);
}

int main(int argc, char** argv)
{
    register_fault_handlers();
    memleak_detect_start();
	ListFunctionTest();
	DListFunctionTest();
    memleak_detect_stop();

    return 0;
}



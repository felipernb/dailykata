#include <stdio.h>

int *slice_array(int *a, int start, int end);
void print_array(int *a) {
	int size = sizeof(a)/sizeof(int);
	int i;
	for (i=0; i < size; i++) {
		printf("%d ", a[i]);
	}
	printf("\n");
}
int *merge_sort(int *a, int start, int end) {
	if (end <= start) return null
	int size = sizeof(a)/sizeof(int);
	int m = size/2;
	int *x = merge_sort(slice_array(a, start, m), 0, m-start);
	int *y = merge_sort(slice_array(a, m+1, end), 0, end-m-1);
	print_array(x);
	print_array(y);
	int merged[size];
	int i, j;
	return merged;
}

int main() {
	int a[5] = {5,4,3,2,1};
	printf("%ld\n", sizeof(a)/sizeof(int));
	int *slice = slice_array(a, 0, 2);
	printf("%ld\n", sizeof(slice)/sizeof(int));
	print_array(slice);
	merge_sort(slice, 0, 2);

}

*slice_array(int *a, int start, int end) {
	int size = end - start;
	int slice[size];
	int i;
	for (i = 0; i <= start-end; i++) {
		slice[i] = a[start+i];
	}
	return slice;
}

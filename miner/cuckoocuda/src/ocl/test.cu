#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include "cuda_profiler_api.h"
#include "cuda_runtime.h"
#include "device_launch_parameters.h"
#include "device_functions.h"
#define n 4

__device__ uint64_t toLong(uint2 a){
	uint64_t l = (uint64_t)a.y << 32 | (uint64_t)a.x;
	return l;
}

__global__ void kernel(uint2 *a, ulonglong4 *b){
	uint2 tmp = {1UL << 31, 1UL << 0};
	for (uint32_t i = 0 ; i < sizeof(uint32_t) * 2; i++ ) {
		printf("%lx ", *((uint8_t*)(((char*)&tmp) + i)));
	}
	printf("\n");
	printf("x  = %x %x\n", tmp.x, tmp.y);
	printf("lx = %lx %lx\n", tmp.x, tmp.y);
	printf("\n");
	printf("1: %lu %lu\n", a[0].x, a[0].y);
	printf("2: %lu %lu\n", a[1].x, a[1].y);
	printf("3: %lu %lu\n", a[2].x, a[2].y);
	printf("4: %lu %lu\n", a[3].x, a[3].y);
	printf("\n");
	printf("1: %u %u\n", a[0].x, a[0].y);
	printf("2: %u %u\n", a[1].x, a[1].y);
	printf("3: %u %u\n", a[2].x, a[2].y);
	printf("4: %u %u\n", a[3].x, a[3].y);
	*b = *(ulonglong4*)(&a[0]);
	printf("%lu %lu %lu %lu\n", toLong(a[0]), toLong(a[1]), toLong(a[2]), toLong(a[3]));

	printf("%llu %llu %llu %llu\n",b->x, b->y, b->z, b->w);

	int xx = 0x1234;
	char cxx = *(char*)&xx;
	if(cxx == 0x12){
		printf("big\n");
	}else{
		printf("little\n");
	}
}

__global__ void kernel2(ulonglong4 *a){
	printf("%lu %lu %lu %lu\n", a[0].x, a[0].y, a[0].z, a[0].w);
}

int main(){
	uint32_t *a = (uint32_t*)malloc(sizeof(uint32_t) * n * 2);
	ulonglong4 b;
	int *deva, *devb;
	cudaMalloc((void**)&deva, sizeof(uint2)*n);
	cudaMalloc((void**)&devb, sizeof(uint2)*n);	
	for(int i = 0; i < 8; i++){
		a[i] = (uint32_t)(i);
//		printf("%llu ", a[i]);
	}
//	printf("\n");
//	printf("%llu\n", sizeof(uint2));
	cudaMemcpy(deva, a, sizeof(uint2)*n, cudaMemcpyHostToDevice);
//	kernel<<<1,1>>>((uint2*)deva, (ulonglong4*)devb);
	kernel2<<<1,1>>>((ulonglong4*)deva);
	cudaMemcpy(&b, devb, sizeof(ulonglong4), cudaMemcpyDeviceToHost);
//	printf("%llu %llu %llu %llu\n", b.x, b.y, b.z, b.w);
	return 0;
}

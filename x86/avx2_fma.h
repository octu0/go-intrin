#include <stdint.h>
#include <string.h>
#include <immintrin.h>

static void imm_fmadd_float32_128(float *out, float *a, float *b, float *c) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 mc = _mm_setr_ps(c[0], c[1], c[2], c[3]);
  __m128 r = _mm_fmadd_ps(ma, mb, mc);
  _mm_store_ps(out, r);
}

static void imm_fmadd_float32_256(float *out, float *a, float *b, float *c) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256 mc = _mm256_setr_ps(c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7]);
  __m256 r = _mm256_fmadd_ps(ma, mb, mc);
  _mm256_store_ps(out, r);
}

static void imm_fmadd_float64_128(double *out, double *a, double *b, double *c) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d mc = _mm_setr_pd(c[0], c[1]);
  __m128 r = _mm_fmadd_pd(ma, mb, mc);
  _mm_store_pd(out, r);
}

static void imm_fmadd_float64_256(double *out, double *a, double *b, double *c) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d mc = _mm256_setr_pd(c[0], c[1], c[2], c[3]);
  __m256 r = _mm256_fmadd_pd(ma, mb, mc);
  _mm256_store_pd(out, r);
}

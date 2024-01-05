#include <stdint.h>
#include <string.h>
#include <immintrin.h>

static void imm_add_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_add_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_sub_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_sub_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_mul_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_mul_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_div_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_div_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_sqrt_float64(double *out, double *a) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d r = _mm256_sqrt_pd(ma);
  _mm256_store_pd(out, r);
}

static void imm_min_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_min_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_max_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_max_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_and_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_and_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_or_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_or_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_xor_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_xor_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_andnot_float64(double *out, double *a, double *b) {
  __m256d ma = _mm256_setr_pd(a[0], a[1], a[2], a[3]);
  __m256d mb = _mm256_setr_pd(b[0], b[1], b[2], b[3]);
  __m256d r = _mm256_andnot_pd(ma, mb);
  _mm256_store_pd(out, r);
}

static void imm_bulk_convert_int32_to_float64(double *out, int32_t *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128i ma = _mm_setr_epi32(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m256 r = _mm256_cvtepi32_pd(ma);
    _mm256_storeu_pd(out + i, r);
  }
}

static void imm_bulk_convert_float64_to_int32(int32_t *out, double *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m256d ma = _mm256_setr_pd(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128i r = _mm256_cvtpd_epi32(ma);
    _mm_storeu_si128((__m128i *) out, r);
  }
}

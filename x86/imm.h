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

static void imm_add_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_add_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_sub_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_sub_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_mul_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_mul_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_div_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_div_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_sqrt_float32(float *out, float *a) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256d r = _mm256_sqrt_ps(ma);
  _mm256_store_ps(out, r);
}

static void imm_rsqrt_float32(float *out, float *a) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256d r = _mm256_rsqrt_ps(ma);
  _mm256_store_ps(out, r);
}

static void imm_rcp_float32(float *out, float *a) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256d r = _mm256_rcp_ps(ma);
  _mm256_store_ps(out, r);
}

static void imm_min_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_min_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_max_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_max_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_and_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_and_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_or_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_or_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_xor_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_xor_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_andnot_float32(float *out, float *a, float *b) {
  __m256 ma = _mm256_setr_ps(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256 mb = _mm256_setr_ps(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256d r = _mm256_andnot_ps(ma, mb);
  _mm256_store_ps(out, r);
}

static void imm_bulk_convert_int32_to_float32(float *out, int32_t *in, int size) {
  for(int i = 0; i < size; i += 8) {
    __m256i ma = _mm256_setr_epi32(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      in[i + 4], in[i + 5], in[i + 6], in[i + 7]
    );
    __m256 r = _mm256_cvtepi32_ps(ma);
    _mm256_storeu_ps(out + i, r);
  }
}

static void imm_bulk_convert_float32_to_int32(int32_t *out, float *in, int size) {
  for(int i = 0; i < size; i += 8) {
    __m256 ma = _mm256_setr_ps(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      in[i + 4], in[i + 5], in[i + 6], in[i + 7]
    );
    __m256i r = _mm256_cvtps_epi32(ma);
    _mm256_storeu_si256((__m256i *)out + i, r);
  }
}

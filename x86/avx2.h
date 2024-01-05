#include <stdint.h>
#include <string.h>
#include <immintrin.h>

static void imm_add_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_add_epi32(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_sub_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_sub_epi32(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_abs_int32(int32_t *out, int32_t *a) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i r = _mm256_abs_epi32(ma);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_and_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_and_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_or_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_or_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_xor_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_xor_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_andnot_int32(int32_t *out, int32_t *a, int32_t *b) {
  __m256i ma = _mm256_setr_epi32(a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
  __m256i mb = _mm256_setr_epi32(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
  __m256i r = _mm256_andnot_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

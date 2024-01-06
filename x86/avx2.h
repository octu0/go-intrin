#include <stdint.h>
#include <string.h>
#include <immintrin.h>

static void imm_add_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_adds_epi8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_sub_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_subs_epi8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_abs_int8(int8_t *out, int8_t *a) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i r = _mm256_abs_epi8(ma);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_min_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_min_epi8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_max_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_max_epi8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_and_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_and_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_or_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_or_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_xor_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_xor_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_andnot_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_andnot_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_add_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_adds_epu8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_sub_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_subs_epu8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_avg_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_avg_epu8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_min_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_min_epu8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_max_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_max_epu8(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_and_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_and_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_or_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_or_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_xor_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_xor_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_andnot_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m256i ma = _mm256_setr_epi8(
    a[0],  a[1],  a[2],  a[3],  a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11], a[12], a[13], a[14], a[15],
    a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23],
    a[24], a[25], a[26], a[27], a[28], a[29], a[30], a[31]
  );
  __m256i mb = _mm256_setr_epi8(
    b[0],  b[1],  b[2],  b[3],  b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11], b[12], b[13], b[14], b[15],
    b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23],
    b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31]
  );
  __m256i r = _mm256_andnot_si256(ma, mb);
  _mm256_store_si256((__m256i *) out, r);
}

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

static void imm_bulk_sum_int8(int8_t *out, int8_t *in, int size) {
  __m256i r = _mm256_setr_epi8(
    in[0],  in[1],  in[2],  in[3],  in[4],  in[5],  in[6],  in[7],
    in[8],  in[9],  in[10], in[11], in[12], in[13], in[14], in[15],
    in[16], in[17], in[18], in[19], in[20], in[21], in[22], in[23],
    in[24], in[25], in[26], in[27], in[28], in[29], in[30], in[31]
  );
  for(int i = 32; i < size; i += 32) {
    __m256i ma = _mm256_setr_epi8(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7],
      in[i + 8],  in[i + 9],  in[i + 10], in[i + 11],
      in[i + 12], in[i + 13], in[i + 14], in[i + 15],
      in[i + 16], in[i + 17], in[i + 18], in[i + 19],
      in[i + 20], in[i + 21], in[i + 22], in[i + 23],
      in[i + 24], in[i + 25], in[i + 26], in[i + 27],
      in[i + 28], in[i + 29], in[i + 30], in[i + 31]
    );
    r = _mm256_adds_epi8(ma, r);
  }
  _mm256_store_si256((__m256i *) out, r);
}

static void imm_bulk_sum_uint8(uint8_t *out, uint8_t *in, int size) {
  __m256i r = _mm256_setr_epi8(
    in[0],  in[1],  in[2],  in[3],  in[4],  in[5],  in[6],  in[7],
    in[8],  in[9],  in[10], in[11], in[12], in[13], in[14], in[15],
    in[16], in[17], in[18], in[19], in[20], in[21], in[22], in[23],
    in[24], in[25], in[26], in[27], in[28], in[29], in[30], in[31]
  );
  for(int i = 32; i < size; i += 32) {
    __m256i ma = _mm256_setr_epi8(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7],
      in[i + 8],  in[i + 9],  in[i + 10], in[i + 11],
      in[i + 12], in[i + 13], in[i + 14], in[i + 15],
      in[i + 16], in[i + 17], in[i + 18], in[i + 19],
      in[i + 20], in[i + 21], in[i + 22], in[i + 23],
      in[i + 24], in[i + 25], in[i + 26], in[i + 27],
      in[i + 28], in[i + 29], in[i + 30], in[i + 31]
    );
    r = _mm256_adds_epu8(ma, r);
  }
  _mm256_store_si256((__m256i *) out, r);
}

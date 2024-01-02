#include <stdint.h>
#include <emmintrin.h>

static void emm_add_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epi8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epi8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_add_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_avg_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  __m128i r = _mm_avg_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_max_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  __m128i r = _mm_max_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_min_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_set_epi8(
    a[15], a[14], a[13], a[12],
    a[11], a[10], a[9],  a[8],
    a[7],  a[6],  a[5],  a[4],
    a[3],  a[2],  a[1],  a[0]
  );
  __m128i mb = _mm_set_epi8(
    b[15], b[14], b[13], b[12],
    b[11], b[10], b[9],  b[8],
    b[7],  b[6] , b[5],  b[4],
    b[3],  b[2],  b[1],  b[0]
  );
  __m128i r = _mm_min_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

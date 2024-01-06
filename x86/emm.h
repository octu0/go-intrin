#include <stdint.h>
#include <string.h>
#include <emmintrin.h>

static void emm_add_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epi8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_int8(int8_t *out, int8_t *a, int8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epi8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_add_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_avg_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  __m128i r = _mm_avg_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_max_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  __m128i r = _mm_max_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_min_uint8(uint8_t *out, uint8_t *a, uint8_t *b) {
  __m128i ma = _mm_setr_epi8(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7],
    a[8],  a[9],  a[10], a[11],
    a[12], a[13], a[14], a[15]
  );
  __m128i mb = _mm_setr_epi8(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7],
    b[8],  b[9],  b[10], b[11],
    b[12], b[13], b[14], b[15]
  );
  __m128i r = _mm_min_epu8(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_add_int16(int16_t *out, int16_t *a, int16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epi16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_int16(int16_t *out, int16_t *a, int16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epi16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_max_int16(int16_t *out, int16_t *a, int16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  __m128i r = _mm_max_epi16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_min_int16(int16_t *out, int16_t *a, int16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  __m128i r = _mm_min_epi16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_madd_int16(int32_t *out, int16_t *a, int16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  __m128i r = _mm_madd_epi16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_add_uint16(uint16_t *out, uint16_t *a, uint16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  // adds = saturation arithmetic
  __m128i r = _mm_adds_epu16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_sub_uint16(uint16_t *out, uint16_t *a, uint16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  // subs = saturation arithmetic
  __m128i r = _mm_subs_epu16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_avg_uint16(uint16_t *out, uint16_t *a, uint16_t *b) {
  __m128i ma = _mm_setr_epi16(
    a[0],  a[1],  a[2],  a[3],
    a[4],  a[5],  a[6],  a[7]
  );
  __m128i mb = _mm_setr_epi16(
    b[0],  b[1],  b[2],  b[3],
    b[4],  b[5],  b[6],  b[7]
  );
  __m128i r = _mm_avg_epu16(ma, mb);
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_add_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_add_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_sub_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_sub_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_mul_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_mul_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_div_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_div_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_sqrt_float64(double *out, double *a) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d r = _mm_sqrt_pd(ma);
  _mm_storeu_pd(out, r);
}

static void emm_max_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_max_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_min_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_min_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_and_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_and_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_or_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_or_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_xor_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_xor_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_andnot_float64(double *out, double *a, double *b) {
  __m128d ma = _mm_setr_pd(a[0], a[1]);
  __m128d mb = _mm_setr_pd(b[0], b[1]);
  __m128d r = _mm_andnot_pd(ma, mb);
  _mm_storeu_pd(out, r);
}

static void emm_bulk_sum_int8(int8_t *out, int8_t *in, int size) {
  __m128i r = _mm_setr_epi8(
    in[0],  in[1],  in[2],  in[3],
    in[4],  in[5],  in[6],  in[7],
    in[8],  in[9],  in[10], in[11],
    in[12], in[13], in[14], in[15]
  );
  for(int i = 16; i < size; i += 16) {
    __m128i ma = _mm_setr_epi8(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7],
      in[i + 8],  in[i + 9],  in[i + 10], in[i + 11],
      in[i + 12], in[i + 13], in[i + 14], in[i + 15]
    );
    r = _mm_adds_epi8(ma, r);
  }
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_bulk_sum_uint8(uint8_t *out, uint8_t *in, int size) {
  __m128i r = _mm_setr_epi8(
    in[0],  in[1],  in[2],  in[3],
    in[4],  in[5],  in[6],  in[7],
    in[8],  in[9],  in[10], in[11],
    in[12], in[13], in[14], in[15]
  );
  for(int i = 16; i < size; i += 16) {
    __m128i ma = _mm_setr_epi8(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7],
      in[i + 8],  in[i + 9],  in[i + 10], in[i + 11],
      in[i + 12], in[i + 13], in[i + 14], in[i + 15]
    );
    r = _mm_adds_epu8(ma, r);
  }
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_bulk_sum_int16(int16_t *out, int16_t *in, int size) {
  __m128i r = _mm_setr_epi16(
    in[0],  in[1],  in[2],  in[3],
    in[4],  in[5],  in[6],  in[7]
  );
  for(int i = 8; i < size; i += 8) {
    __m128i ma = _mm_setr_epi16(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7]
    );
    r = _mm_adds_epi16(ma, r);
  }
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_bulk_sum_uint16(uint16_t *out, uint16_t *in, int size) {
  __m128i r = _mm_setr_epi16(
    in[0],  in[1],  in[2],  in[3],
    in[4],  in[5],  in[6],  in[7]
  );
  for(int i = 8; i < size; i += 8) {
    __m128i ma = _mm_setr_epi16(
      in[i + 0],  in[i + 1],  in[i + 2],  in[i + 3],
      in[i + 4],  in[i + 5],  in[i + 6],  in[i + 7]
    );
    r = _mm_adds_epu16(ma, r);
  }
  _mm_storeu_si128((__m128i *) out, r);
}

static void emm_bulk_convert_int32_to_float32(float *out, int32_t *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128i ma = _mm_setr_epi32(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_cvtepi32_ps(ma);
    _mm_stream_ps(out + i, r);
  }
}

static void emm_bulk_convert_float32_to_int32(int32_t *out, float *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128i r = _mm_cvtps_epi32(ma);
    _mm_stream_si128((__m128i *) out + i, r);
  }
}

static void emm_bulk_convert_int32_to_float64(double *out, int32_t *in, int size) {
  for(int i = 0; i < size; i += 2) {
    __m64 ma = _mm_setr_pi32(in[i + 0], in[i + 1]);
    __m128d r = _mm_cvtpi32_pd(ma);
    _mm_store_pd(out + i, r);
  }
}

static void emm_bulk_convert_float64_to_int32(int32_t *out, double *in, int size) {
  for(int i = 0; i < size; i += 2) {
    __m128d ma = _mm_setr_pd(in[i + 0], in[i + 1]);
    __m64 r = _mm_cvtpd_pi32(ma);
    memcpy(out + i, &r, sizeof(int32_t) * 2);
  }
}

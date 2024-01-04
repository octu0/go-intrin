#include <stdint.h>
#include <string.h>
#include <xmmintrin.h>

static void xmm_add(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_add_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_sub(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_sub_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_mul(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_mul_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_div(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_div_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_sqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_sqrt_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_rsqrt(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rsqrt_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_rcp(float *out, float *a) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 r = _mm_rcp_ps(ma);
  _mm_store_ps(out, r);
}

static void xmm_min(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_min_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_max(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_max_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_and(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_and_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_or(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_or_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_xor(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_xor_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_andnot(float *out, float *a, float *b) {
  __m128 ma = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  __m128 mb = _mm_setr_ps(b[0], b[1], b[2], b[3]);
  __m128 r = _mm_andnot_ps(ma, mb);
  _mm_store_ps(out, r);
}

static void xmm_bulk_convert_int8_to_float32(float *out, int8_t *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m64 ma = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m128 r = _mm_cvtpi8_ps(ma);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_convert_uint8_to_float32(float *out, uint8_t *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m64 ma = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m128 r = _mm_cvtpu8_ps(ma);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_convert_float32_to_int8(int8_t *out, float *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m64 r = _mm_cvtps_pi8(ma);
    memcpy(out + i, &r, sizeof(char) * 4);
  }
}

static void xmm_bulk_convert_float32_to_uint8(uint8_t *out, float *in, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m64 r = _mm_cvtps_pi8(ma);
    memcpy(out + i, &r, sizeof(unsigned char) * 4);
  }
}

static void xmm_bulk_sum(float *out, float *a, int size) {
  __m128 r = _mm_setr_ps(a[0], a[1], a[2], a[3]);
  for(int i = 4; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(a[i + 0], a[i + 1], a[i + 2], a[i + 3]);
    r = _mm_add_ps(ma, r);
  }
  _mm_store_ps(out, r);
}

static void xmm_bulk_sum2(float *out, float *a, float *b, int size) {
  for(int i = 0; i < size; i += 4) {
    __m128 ma = _mm_setr_ps(a[i + 0], a[i + 1], a[i + 2], a[i + 3]);
    __m128 mb = _mm_setr_ps(b[i + 0], b[i + 1], b[i + 2], b[i + 3]);
    __m128 r = _mm_add_ps(ma, mb);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_add(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  for(int i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_add_ps(min, mbase);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_sub(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  for(int i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_sub_ps(min, mbase);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_mul(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  for(int i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_mul_ps(min, mbase);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_bulk_div(float *out, float *base, float *in, int size) {
  __m128 mbase = _mm_setr_ps(base[0], base[1], base[2], base[3]);
  for(int i = 0; i < size; i += 4) {
    __m128 min = _mm_setr_ps(in[i + 0], in[i + 1], in[i + 2], in[i + 3]);
    __m128 r = _mm_div_ps(min, mbase);
    _mm_stream_ps(out + i, r);
  }
}

static void xmm_tile4x4_sum(float *out, float *in, int size) {
  int i = 0;
  int p = 0;
  for(; i < size;) {
    __m128 r = _mm_setr_ps(in[i + 0 + 0], in[i + 4 + 0], in[i + 8 + 0], in[i + 12 + 0]); // R
    __m128 g = _mm_setr_ps(in[i + 0 + 1], in[i + 4 + 1], in[i + 8 + 1], in[i + 12 + 1]); // G
    __m128 b = _mm_setr_ps(in[i + 0 + 2], in[i + 4 + 2], in[i + 8 + 2], in[i + 12 + 2]); // B
    __m128 a = _mm_setr_ps(in[i + 0 + 3], in[i + 4 + 3], in[i + 8 + 3], in[i + 12 + 3]); // A
    
    __m128 rg = _mm_add_ps(r, g);
    __m128 ba = _mm_add_ps(b, a);
    __m128 gray = _mm_add_ps(rg, ba);
    _mm_stream_ps(out + p, gray);

    i += 16;
    p += 4;
  }
}

static void xmm_grayscale(uint8_t *out, uint8_t *in, int size) {
  __m128 bt709 = _mm_setr_ps(0.2126f, 0.7152f, 0.0722f, 0.0f);
  uint8_t gray[8];
  for(int i = 0; i < size; i += 16) {
    __m64 m1 = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m64 m2 = _mm_setr_pi8(
      in[i + 4], in[i + 5], in[i + 6], in[i + 7],
      0, 0, 0, 0
    );
    __m64 m3 = _mm_setr_pi8(
      in[i + 8], in[i + 9], in[i + 10], in[i + 11],
      0, 0, 0, 0
    );
    __m64 m4 = _mm_setr_pi8(
      in[i + 12], in[i + 13], in[i + 14], in[i + 15],
      0, 0, 0, 0
    );
    __m128 rgba1 = _mm_mul_ps(_mm_cvtpu8_ps(m1), bt701);
    __m128 rgba2 = _mm_mul_ps(_mm_cvtpu8_ps(m2), bt701);
    __m128 rgba3 = _mm_mul_ps(_mm_cvtpu8_ps(m3), bt701);
    __m128 rgba4 = _mm_mul_ps(_mm_cvtpu8_ps(m4), bt701);

    __m128 r = _mm_setr_ps(rgba1[0], rgba2[0], rgba3[0], rgba4[0]); // R
    __m128 g = _mm_setr_ps(rgba1[1], rgba2[1], rgba3[1], rgba4[1]); // G
    __m128 b = _mm_setr_ps(rgba1[2], rgba2[2], rgba3[2], rgba4[2]); // B

    // gray = [rgba1,rgba2,rgba3,rgba4]
    __m128 gray_float = _mm_add_ps(_mm_add_ps(r, g), b);

    __m64 gray_u8 = _mm_cvtps_pi8(gray_float);
    memcpy(&gray, &gray_u8, sizeof(__m64));

    uint8_t tmp[16] = {
      gray[0], gray[0], gray[0], 255,
      gray[1], gray[1], gray[1], 255,
      gray[2], gray[2], gray[2], 255,
      gray[3], gray[3], gray[3], 255
    };
    memcpy(out + i, &tmp, 16);
  }
}

static void xmm_grayscale_in(uint8_t *out, uint8_t *in, int size) {
  float *fin = (float *) malloc(sizeof(float) * size);
  xmm_bulk_convert_uint8_to_float32(fin, in, size);
  float filter[4] = {0.2126f, 0.7152f, 0.0722f, 0.0f};
  xmm_bulk_mul(fin, filter, fin, size);
  float *tile = (float *) malloc(sizeof(float) * (size/4));
  xmm_tile4x4_sum(tile, fin, size);
  uint8_t *gray = (uint8_t *) malloc(sizeof(uint8_t) * (size/4));
  xmm_bulk_convert_float32_to_uint8(gray, tile, (size/4));
  free(fin);
  free(tile);

  int gray_offset = 0;
  for(int i = 0; i < size; i += 4) {
    __m64 ma = _mm_setr_pi8(
      gray[gray_offset], gray[gray_offset], gray[gray_offset], -1,
      0, 0, 0, 0
    );
    memcpy(out + i, &ma, sizeof(uint8_t) * 4);
    gray_offset += 1;
  }
  free(gray);
}

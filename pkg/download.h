/** -*- mode: C; -*-
 * download.h
 * Copyright (C) 2021 Nyah Check
 *
 **/

#ifndef DOWNLOAD_H
#define DOWNLOAD_H
#include <stdint.h>
#include <inttypes.h>

// result type
typedef struct {
  int32_t Ok;
  int32_t Err;
} result_t;

result_t download(const char *url, const char *path);

#endif

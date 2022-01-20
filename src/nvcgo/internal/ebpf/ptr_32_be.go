//go:build armbe || mips || mips64p32
// +build armbe mips mips64p32

/*
 * Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file is inspired by code in the (MIT licensed) cilium eBPF package at:
// https://github.com/cilium/ebpf/blob/b38550c6f15200e3798c695890f699001b97e229/internal/ptr_32_be.go

package ebpf

import (
	"unsafe"
)

// Pointer wraps an unsafe.Pointer to be 64bit to
// conform to the syscall specification.
type Pointer struct {
	pad uint32
	ptr unsafe.Pointer
}
// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/mpi.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fmpi_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fmpi_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fmpi_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fmpi_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fmpi_2eproto();
namespace flyteidl {
namespace plugins {
class DistributedMPITrainingTask;
class DistributedMPITrainingTaskDefaultTypeInternal;
extern DistributedMPITrainingTaskDefaultTypeInternal _DistributedMPITrainingTask_default_instance_;
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::DistributedMPITrainingTask* Arena::CreateMaybeMessage<::flyteidl::plugins::DistributedMPITrainingTask>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {

// ===================================================================

class DistributedMPITrainingTask final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.DistributedMPITrainingTask) */ {
 public:
  DistributedMPITrainingTask();
  virtual ~DistributedMPITrainingTask();

  DistributedMPITrainingTask(const DistributedMPITrainingTask& from);

  inline DistributedMPITrainingTask& operator=(const DistributedMPITrainingTask& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DistributedMPITrainingTask(DistributedMPITrainingTask&& from) noexcept
    : DistributedMPITrainingTask() {
    *this = ::std::move(from);
  }

  inline DistributedMPITrainingTask& operator=(DistributedMPITrainingTask&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const DistributedMPITrainingTask& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DistributedMPITrainingTask* internal_default_instance() {
    return reinterpret_cast<const DistributedMPITrainingTask*>(
               &_DistributedMPITrainingTask_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(DistributedMPITrainingTask* other);
  friend void swap(DistributedMPITrainingTask& a, DistributedMPITrainingTask& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DistributedMPITrainingTask* New() const final {
    return CreateMaybeMessage<DistributedMPITrainingTask>(nullptr);
  }

  DistributedMPITrainingTask* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<DistributedMPITrainingTask>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const DistributedMPITrainingTask& from);
  void MergeFrom(const DistributedMPITrainingTask& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(DistributedMPITrainingTask* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // int32 num_workers = 1;
  void clear_num_workers();
  static const int kNumWorkersFieldNumber = 1;
  ::google::protobuf::int32 num_workers() const;
  void set_num_workers(::google::protobuf::int32 value);

  // int32 num_launcher_replicas = 2;
  void clear_num_launcher_replicas();
  static const int kNumLauncherReplicasFieldNumber = 2;
  ::google::protobuf::int32 num_launcher_replicas() const;
  void set_num_launcher_replicas(::google::protobuf::int32 value);

  // int32 slots = 3;
  void clear_slots();
  static const int kSlotsFieldNumber = 3;
  ::google::protobuf::int32 slots() const;
  void set_slots(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.DistributedMPITrainingTask)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::int32 num_workers_;
  ::google::protobuf::int32 num_launcher_replicas_;
  ::google::protobuf::int32 slots_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fmpi_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// DistributedMPITrainingTask

// int32 num_workers = 1;
inline void DistributedMPITrainingTask::clear_num_workers() {
  num_workers_ = 0;
}
inline ::google::protobuf::int32 DistributedMPITrainingTask::num_workers() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedMPITrainingTask.num_workers)
  return num_workers_;
}
inline void DistributedMPITrainingTask::set_num_workers(::google::protobuf::int32 value) {
  
  num_workers_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedMPITrainingTask.num_workers)
}

// int32 num_launcher_replicas = 2;
inline void DistributedMPITrainingTask::clear_num_launcher_replicas() {
  num_launcher_replicas_ = 0;
}
inline ::google::protobuf::int32 DistributedMPITrainingTask::num_launcher_replicas() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedMPITrainingTask.num_launcher_replicas)
  return num_launcher_replicas_;
}
inline void DistributedMPITrainingTask::set_num_launcher_replicas(::google::protobuf::int32 value) {
  
  num_launcher_replicas_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedMPITrainingTask.num_launcher_replicas)
}

// int32 slots = 3;
inline void DistributedMPITrainingTask::clear_slots() {
  slots_ = 0;
}
inline ::google::protobuf::int32 DistributedMPITrainingTask::slots() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.DistributedMPITrainingTask.slots)
  return slots_;
}
inline void DistributedMPITrainingTask::set_slots(::google::protobuf::int32 value) {
  
  slots_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.DistributedMPITrainingTask.slots)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace plugins
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fmpi_2eproto

package iavl

import (
	iavlOrig "github.com/tendermint/iavl/iavl"
)

var (
	NewMutableTree = iavlOrig.NewMutableTree
	NewIAVLValueOp = iavlOrig.NewIAVLValueOp
	NewIAVLAbsenceOp = iavlOrig.NewIAVLAbsenceOp
	IAVLValueOpDecoder = iavlOrig.IAVLValueOpDecoder
	IAVLAbsenceOpDecoder = iavlOrig.IAVLAbsenceOpDecoder

	ErrVersionDoesNotExist = iavlOrig.ErrVersionDoesNotExist
)


const (
	ProofOpIAVLValue = iavlOrig.ProofOpIAVLValue
	ProofOpIAVLAbsence = iavlOrig.ProofOpIAVLAbsence
)

type (
	RangeProof = iavlOrig.RangeProof
	MutableTree = iavlOrig.MutableTree
	ImmutableTree = iavlOrig.ImmutableTree
)

/*
func NewMutableTree(db dbm.DB, cacheSize int) *MutableTree {

func NewIAVLValueOp(key []byte, proof *RangeProof) IAVLValueOp {
const ProofOpIAVLValue = "iavl:v"
type IAVLValueOp struct {
	// Encoded in ProofOp.Key.
	key []byte

	// To encode in ProofOp.Data.
	// Proof is nil for an empty tree.
	// The hash of an empty tree is nil.
	Proof *RangeProof `json:"proof"`
}
//type ProofOperator interface {
//    Run([][]byte) ([][]byte, error)
//    GetKey() []byte
//    ProofOp() ProofOp
//}
var _ merkle.ProofOperator = IAVLValueOp{}
func NewIAVLValueOp(key []byte, proof *RangeProof) IAVLValueOp {
func IAVLValueOpDecoder(pop merkle.ProofOp) (merkle.ProofOperator, error) {
func (op IAVLValueOp) ProofOp() merkle.ProofOp {
func (op IAVLValueOp) String() string {
func (op IAVLValueOp) Run(args [][]byte) ([][]byte, error) {
func (op IAVLValueOp) GetKey() []byte {

const ProofOpIAVLAbsence = "iavl:a"
type IAVLAbsenceOp struct {
	// Encoded in ProofOp.Key.
	key []byte

	// To encode in ProofOp.Data.
	// Proof is nil for an empty tree.
	// The hash of an empty tree is nil.
	Proof *RangeProof `json:"proof"`
}
var _ merkle.ProofOperator = IAVLAbsenceOp{}
func NewIAVLAbsenceOp(key []byte, proof *RangeProof) IAVLAbsenceOp {
func IAVLAbsenceOpDecoder(pop merkle.ProofOp) (merkle.ProofOperator, error) {
func (op IAVLAbsenceOp) ProofOp() merkle.ProofOp {
func (op IAVLAbsenceOp) String() string {
func (op IAVLAbsenceOp) Run(args [][]byte) ([][]byte, error) {
func (op IAVLAbsenceOp) GetKey() []byte {

func IAVLValueOpDecoder(pop merkle.ProofOp) (merkle.ProofOperator, error) {
func IAVLAbsenceOpDecoder(pop merkle.ProofOp) (merkle.ProofOperator, error) {

var ErrVersionDoesNotExist = fmt.Errorf("version does not exist")

type RangeProof struct {
	// You don't need the right path because
	// it can be derived from what we have.
	LeftPath   PathToLeaf      `json:"left_path"`
	InnerNodes []PathToLeaf    `json:"inner_nodes"`
	Leaves     []proofLeafNode `json:"leaves"`

	// memoize
	rootVerified bool
	rootHash     []byte // valid iff rootVerified is true
	treeEnd      bool   // valid iff rootVerified is true

}
func (proof *RangeProof) Keys() (keys [][]byte) {
func (proof *RangeProof) String() string {
func (proof *RangeProof) StringIndented(indent string) string {
func (proof *RangeProof) LeftIndex() int64 {
func (proof *RangeProof) VerifyItem(key, value []byte) error {
func (proof *RangeProof) VerifyAbsence(key []byte) error {
func (proof *RangeProof) Verify(root []byte) error {
func (proof *RangeProof) ComputeRootHash() []byte {

func (t *ImmutableTree) GetWithProof(key []byte) (value []byte, proof *RangeProof, err error) {
func (t *ImmutableTree) GetRangeWithProof(startKey []byte, endKey []byte, limit int) (keys, values [][]byte, proof *RangeProof, err error) {
func (tree *MutableTree) GetVersionedWithProof(key []byte, version int64) ([]byte, *RangeProof, error) {
func (tree *MutableTree) GetVersionedRangeWithProof(startKey, endKey []byte, limit int, version int64) (

func (tree *MutableTree) AvailableVersions() []int {
func (tree *MutableTree) DeleteVersion(version int64) error {
func (tree *MutableTree) GetImmutable(version int64) (*ImmutableTree, error) {
func (tree *MutableTree) GetVersioned(key []byte, version int64) (
func (tree *MutableTree) Hash() []byte {
func (tree *MutableTree) IsEmpty() bool {
func (tree *MutableTree) LazyLoadVersion(targetVersion int64) (int64, error) {
func (tree *MutableTree) Load() (int64, error) {
func (tree *MutableTree) LoadVersion(targetVersion int64) (int64, error) {
func (tree *MutableTree) LoadVersionForOverwriting(targetVersion int64) (int64, error) {
func (tree *MutableTree) Remove(key []byte) ([]byte, bool) {
func (tree *MutableTree) Rollback() {
func (tree *MutableTree) SaveVersion() ([]byte, int64, error) {
func (tree *MutableTree) Set(key, value []byte) bool {
func (tree *MutableTree) String() string {
func (tree *MutableTree) VersionExists(version int64) bool {
func (tree *MutableTree) WorkingHash() []byte {

func (t *ImmutableTree) Get(key []byte) (index int64, value []byte) {
func (t *ImmutableTree) GetByIndex(index int64) (key []byte, value []byte) {
func (t *ImmutableTree) Has(key []byte) bool {
func (t *ImmutableTree) Hash() []byte {
func (t *ImmutableTree) Height() int8 {
func (t *ImmutableTree) Iterate(fn func(key []byte, value []byte) bool) (stopped bool) {
func (t *ImmutableTree) IterateRange(start, end []byte, ascending bool, fn func(key []byte, value []byte) bool) (stopped bool) {
func (t *ImmutableTree) IterateRangeInclusive(start, end []byte, ascending bool, fn func(key, value []byte, version int64) bool) (stopped bool) {
func (t *ImmutableTree) RenderShape(indent string, encoder NodeEncoder) []string {
func (t *ImmutableTree) Size() int64 {
func (t *ImmutableTree) String() string {
func (t *ImmutableTree) Version() int64 {
*/

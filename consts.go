package gexto

const BAD_INO = 1
const ROOT_INO = 2
const USR_QUOTA_INO = 3
const GRP_QUOTA_INO = 4
const BOOT_LOADER_INO = 5
const UNDEL_DIR_INO = 6
const RESIZE_INO = 7
const JOURNAL_INO = 8

const BG_INODE_UNINIT = 0x0001
const BG_BLOCK_UNINIT = 0x0002
const BG_INODE_ZEROED = 0x0004

// Inode flags
const SECRM_FL = 0x00000001
const UNRM_FL = 0x00000002
const COMPR_FL = 0x00000004
const SYNC_FL = 0x00000008
const IMMUTABLE_FL = 0x00000010
const APPEND_FL = 0x00000020
const NODUMP_FL = 0x00000040
const NOATIME_FL = 0x00000080
const DIRTY_FL = 0x00000100
const COMPRBLK_FL = 0x00000200
const NOCOMPR_FL = 0x00000400
const ENCRYPT_FL = 0x00000800
const INDEX_FL = 0x00001000
const IMAGIC_FL = 0x00002000
const JOURNAL_DATA_FL = 0x00004000
const NOTAIL_FL = 0x00008000
const DIRSYNC_FL = 0x00010000
const TOPDIR_FL = 0x00020000
const HUGE_FILE_FL = 0x00040000
const EXTENTS_FL = 0x00080000
const EA_INODE_FL = 0x00200000
const EOFBLOCKS_FL = 0x00400000
const INLINE_DATA_FL = 0x10000000
const PROJINHERIT_FL = 0x20000000
const RESERVED_FL = 0x80000000

const FL_USER_VISIBLE = 0x304BDFFF
const FL_USER_MODIFIABLE = 0x204BC0FF

const (
	EXT4_INODE_SECRM        = 0
	EXT4_INODE_UNRM         = 1
	EXT4_INODE_COMPR        = 2
	EXT4_INODE_SYNC         = 3
	EXT4_INODE_IMMUTABLE    = 4
	EXT4_INODE_APPEND       = 5
	EXT4_INODE_NODUMP       = 6
	EXT4_INODE_NOATIME      = 7
	EXT4_INODE_DIRTY        = 8
	EXT4_INODE_COMPRBLK     = 9
	EXT4_INODE_NOCOMPR      = 10
	EXT4_INODE_ENCRYPT      = 11
	EXT4_INODE_INDEX        = 12
	EXT4_INODE_IMAGIC       = 13
	EXT4_INODE_JOURNAL_DATA = 14
	EXT4_INODE_NOTAIL       = 15
	EXT4_INODE_DIRSYNC      = 16
	EXT4_INODE_TOPDIR       = 17
	EXT4_INODE_HUGE_FILE    = 18
	EXT4_INODE_EXTENTS      = 19
	EXT4_INODE_EA_INODE     = 21
	EXT4_INODE_EOFBLOCKS    = 22
	EXT4_INODE_INLINE_DATA  = 28
	EXT4_INODE_PROJINHERIT  = 29
	EXT4_INODE_RESERVED     = 31
)

const FEATURE_COMPAT_DIR_PREALLOC = 0x0001
const FEATURE_COMPAT_IMAGIC_INODES = 0x0002
const FEATURE_COMPAT_HAS_JOURNAL = 0x0004
const FEATURE_COMPAT_EXT_ATTR = 0x0008
const FEATURE_COMPAT_RESIZE_INODE = 0x0010
const FEATURE_COMPAT_DIR_INDEX = 0x0020
const FEATURE_COMPAT_SPARSE_SUPER2 = 0x0200

const FEATURE_RO_COMPAT_SPARSE_SUPER = 0x0001
const FEATURE_RO_COMPAT_LARGE_FILE = 0x0002
const FEATURE_RO_COMPAT_BTREE_DIR = 0x0004
const FEATURE_RO_COMPAT_HUGE_FILE = 0x0008
const FEATURE_RO_COMPAT_GDT_CSUM = 0x0010
const FEATURE_RO_COMPAT_DIR_NLINK = 0x0020
const FEATURE_RO_COMPAT_EXTRA_ISIZE = 0x0040
const FEATURE_RO_COMPAT_QUOTA = 0x0100
const FEATURE_RO_COMPAT_BIGALLOC = 0x0200
const FEATURE_RO_COMPAT_METADATA_CSUM = 0x0400
const FEATURE_RO_COMPAT_READONLY = 0x1000
const FEATURE_RO_COMPAT_PROJECT = 0x2000

const FEATURE_INCOMPAT_COMPRESSION = 0x0001
const FEATURE_INCOMPAT_FILETYPE = 0x0002
const FEATURE_INCOMPAT_RECOVER = 0x0004
const FEATURE_INCOMPAT_JOURNAL_DEV = 0x0008
const FEATURE_INCOMPAT_META_BG = 0x0010
const FEATURE_INCOMPAT_EXTENTS = 0x0040
const FEATURE_INCOMPAT_64BIT = 0x0080
const FEATURE_INCOMPAT_MMP = 0x0100
const FEATURE_INCOMPAT_FLEX_BG = 0x0200
const FEATURE_INCOMPAT_EA_INODE = 0x0400
const FEATURE_INCOMPAT_DIRDATA = 0x1000
const FEATURE_INCOMPAT_CSUM_SEED = 0x2000
const FEATURE_INCOMPAT_LARGEDIR = 0x4000
const FEATURE_INCOMPAT_INLINE_DATA = 0x8000
const FEATURE_INCOMPAT_ENCRYPT = 0x10000

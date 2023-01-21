# ArchiveBox Pinboard.in Tag Transformer

Utility to convert Pinboard.in entries and tags into a format for importing into ArchiveBox v0.6.3.

The goal is to support multi-word tags with spaces in ArchiveBox. This is not supported by Pinboard.in, where the space character is
used as a separator between tags. To input new Pinboard.in tags the hyphen character is used as a separator. This is
reasonably efficient to enter on mobile keyboards.

    Pinboard Tag: right-to-repair

Importing the Pinboard.in RSS feed into ArchiveBox is first transformed to map the hyphen separated tags
into a format using space characters.

The ArchiveBox tag definition would be:

    Name: Right To Repair
    Slug: right-to-repair

The goal is that most/all of the tags can be edited as part of the process of adding a new entry
to Pinboard.in, rather than adding tags once imported into ArchiveBox.

## Transformer Implementation

The tag transformer is implemented in Go.

## Importing Pinboard.in RSS URLs and Tags

The normal ArchiveBox v0.6.3 import process for Pinboard.in RSS feeds is followed, with the addition
of transforming the list of tags in the <dc:subject>....</dc:subject> XML elements.

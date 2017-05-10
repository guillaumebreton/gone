class Gone < Formula
  desc ""
  homepage ""
  url "https://github.com/guillaumebreton/gone/releases/download/2.3.3/gone_Darwin_x86_64.tar.gz"
  version "2.3.3"
  sha256 "f6d7942f3c67987c1763426196acfd9f81cbc8f9fcb81ae8fe18507ca4577334"

  def install
    bin.install "gone"
  end
end

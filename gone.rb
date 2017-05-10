class Gone < Formula
  desc ""
  homepage ""
  url "https://github.com/guillaumebreton/gone/releases/download/2.3.6/gone_Darwin_x86_64.tar.gz"
  version "2.3.6"
  sha256 "7b0728e137d86659d8230b0f1c4f9e8cb3af8d1a67d4477f0a74a23f0ca7a223"

  def install
    bin.install "gone"
  end
end

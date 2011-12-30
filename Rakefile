require 'rake/clean'
require 'fileutils'

BUILD_DIR = Dir.getwd + '/out'
GO = FileList['**/*.go']
GO_6  = GO.ext('6')

CLOBBER.include(GO_6)

rule '.6' => '.go' do |t|
    if not File.directory? BUILD_DIR then
        Dir.mkdir(BUILD_DIR)
    end
    sh "6g -o #{BUILD_DIR + "/" + File.basename(t.name)} #{t.source}" 
end

task :default => [:compile] do
   sh "6l -o #{BUILD_DIR}/livereload #{FileList[BUILD_DIR + '/**/*.6']}"
end

task :compile => GO.ext('6')

task :clean do |t|
    FileUtils.rm_rf BUILD_DIR
end

task :watch => [:compile] do
    require 'rubygems'
    require 'fssm'

    def rebuild
        sh 'rake'
        puts "    OK"
    rescue
        nil
    end

    begin
        FSSM.monitor(nil, ['**/*.go']) do
            update { rebuild }
            delete { rebuild }
            create { rebuild }
        end
    rescue FSSM::CallbackError => e
        Process.exit
    end
end

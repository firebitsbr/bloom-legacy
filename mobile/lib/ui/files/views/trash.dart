import 'package:bloom/ui/files/widgets/drawer.dart';
import 'package:bloom/libs/filesize.dart';
import 'package:flutter/material.dart';

class FilesTrashView extends StatefulWidget {
  const FilesTrashView();

  @override
  _FilesTrashState createState() => _FilesTrashState();
}

class _FilesTrashState extends State<FilesTrashView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      drawer: const FilesDrawer(),
      appBar: AppBar(
        title: const Text('Files'),
      ),
      body: _buildBody(),
      floatingActionButton: _buildFloatingActionButton(),
    );
  }

  Widget _buildBody() {
    final List<File> files = File.getFiles();

    return Container(
      child: ListView.builder(
          itemCount: files.length,
          itemBuilder: (BuildContext context, int index) {
            final File file = files[index];

            return ListTile(
              leading: Icon(Icons.movie),
              title: Text(file.name),
              subtitle: Text(filesize(file.size)),
              trailing: Icon(Icons.more_vert),
            );
          }),
    );
  }

  FloatingActionButton _buildFloatingActionButton() {
    return FloatingActionButton(
      onPressed: () => debugPrint('Add download pressed'),
      child: Icon(Icons.add),
      backgroundColor: Colors.red,
    );
  }
}

class File {
  File({this.name, this.size}) {
    updatedAt = DateTime.now();
  }

  DateTime updatedAt;
  String name;
  int size;

  static List<File> getFiles() {
    return <File>[
      File(name: 'My super clip', size: 42000000),
    ];
  }
}
